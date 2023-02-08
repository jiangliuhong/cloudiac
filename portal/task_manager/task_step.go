// Copyright (c) 2015-2023 CloudJ Technology Co., Ltd.

package task_manager

import (
	"cloudiac/common"
	"cloudiac/portal/models"
	"cloudiac/portal/services"
	"cloudiac/runner"
	"context"

	"github.com/pkg/errors"
)

// 任务的流程步骤执行结束后的操作
func (m *TaskManager) runTaskStepsDoneActions(ctx context.Context, taskId models.Id) error {
	logger := m.logger.WithField("func", "runTaskStepsDoneActions").WithField("taskId", taskId)

	task, er := services.GetTaskById(m.db, taskId)
	if er != nil {
		return errors.Wrapf(er, "query task %s", taskId)
	}

	runTaskReq, err := buildRunTaskReq(m.db, *task)
	if err != nil {
		return err
	}

	currStep, err := services.GetTaskStep(m.db, task.Id, task.CurrStep)
	if err != nil {
		return errors.Wrapf(err, "get task current step")
	}

	// 执行 callback 步骤
	m.taskStepDoneCallback(ctx, task, currStep, *runTaskReq)

	// 未执行过 apply 操作的任务进行信息采集
	if task.IsEffectTask() && !currStep.IsRejected() && !task.Applied {
		// 执行信息采集步骤
		logger.Infof("run task collect step")
		if err := m.runTaskStep(ctx, *runTaskReq, task, &models.TaskStep{
			PipelineStep: models.PipelineStep{
				Type: models.TaskStepCollect,
			},
			OrgId:     task.OrgId,
			ProjectId: task.ProjectId,
			EnvId:     task.EnvId,
			TaskId:    task.Id,
			Index:     common.CollectTaskStepIndex,
			Status:    models.TaskStepPending,
		}); err != nil {
			logger.Errorf("run task collect step: %v", err)
		}
	}

	return nil
}

func (m *TaskManager) taskStepDoneCallback(ctx context.Context, task *models.Task, currStep *models.TaskStep, runTaskReq runner.RunTaskReq) {

	logger := m.logger.
		WithField("func", "taskStepDoneCallback").
		WithField("taskId", task.Id).
		WithField("step", currStep)

	taskLastStep, err := services.GetTaskLastStep(m.db, task.Id)
	if err != nil {
		logger.Errorf("get task last step: %v", err)
		return
	}

	callbackSteps := make([]models.PipelineStep, 0)
	if currStep.IsSuccess() && task.Flow.OnSuccess != nil {
		step := *task.Flow.OnSuccess
		if step.Name == "" {
			step.Name = "onSuccess"
		}
		callbackSteps = append(callbackSteps, step)
	}

	if currStep.IsFail() && task.Flow.OnFail != nil {
		step := *task.Flow.OnFail
		if step.Name == "" {
			step.Name = "onFail"
		}
		callbackSteps = append(callbackSteps, step)
	}

	newStepIndex := taskLastStep.Index + 1
	for _, pipelineStep := range callbackSteps {
		step, err := services.CreateTaskCallbackStep(m.db, *task, pipelineStep, newStepIndex)
		if err != nil {
			logger.Errorf("create task callback step: %v", err)
			continue
		}

		newStepIndex += 1
		logger.Infof("run task callback step: %s", step.Id)
		if err := m.runTaskStep(ctx, runTaskReq, task, step); err != nil {
			logger.Infof("run task callback step: %v", err)
		}
	}
}
