package usecase

import (
	"context"
	models "inter/internal"
	"sort"
	"sync"
	"time"
)

type UseCase struct {
	InMemory            map[int]*models.Task
	CountOfGoingRoutine int
	Queue               []int
	ctx                 *context.Context
	LengthQueue         int
	Mutex               sync.RWMutex
	Current             int
}

func InitUseCase(countRoutine int, ctx context.Context, closeChan chan bool) models.UseCase {
	useCase := UseCase{}
	useCase.LengthQueue = 0
	useCase.Current = 0
	useCase.CountOfGoingRoutine = countRoutine
	useCase.InMemory = make(map[int]*models.Task)
	go useCase.handleStarter(ctx, closeChan)
	return &useCase
}

func (useCase *UseCase) Create(task models.Task) {
	operator := func(task *models.Task) {
		useCase.Mutex.Lock()
		task.Status = "In queue"
		task.Place = useCase.Current
		useCase.Current++
		task.CreationTime = time.Now().Format("15:04:05 02.01.2006")
		useCase.InMemory[task.Place] = task
		useCase.Queue = append(useCase.Queue, task.Place)
		useCase.LengthQueue++
		useCase.Mutex.Unlock()
	}
	operator(&task)
}

func (useCase *UseCase) handleStarter(ctx context.Context, closeChan chan bool) {
	var mutex sync.Mutex
	for {
		select {
		case <-ctx.Done():
			closeChan <- true
			return
		default:
			if useCase.LengthQueue >= useCase.CountOfGoingRoutine {
				go useCase.Handle(&mutex)
				useCase.LengthQueue -= useCase.CountOfGoingRoutine
			}
		}
	}
}

func (useCase *UseCase) Get() []models.Task {
	var result []models.Task
	for key, _ := range useCase.InMemory {
		result = append(result, *useCase.InMemory[key])
	}
	sort.Slice(result, func(i int, j int) bool {
		return result[i].Place < result[j].Place
	})
	return result
}

func (useCase *UseCase) Handle(mut *sync.Mutex) {
	mut.Lock()
	var wg sync.WaitGroup
	sort.Slice(useCase.Queue, func(i, j int) bool {
		return useCase.Queue[i] < useCase.Queue[j]
	})

	for i := 0; i < useCase.CountOfGoingRoutine; i++ {
		go useCase.onGoingOperation(useCase.InMemory[useCase.Queue[i]], &wg)
		wg.Add(1)
	}
	wg.Wait()
	var deleteMutex sync.Mutex
	useCase.Deleter(&deleteMutex)
	mut.Unlock()
}

func (useCase *UseCase) Deleter(mutex *sync.Mutex) {
	mutex.Lock()
	if len(useCase.Queue) == useCase.CountOfGoingRoutine {
		useCase.Queue = nil
		return
	} else if len(useCase.Queue) < useCase.CountOfGoingRoutine {
		mutex.Unlock()
		return
	}
	copy(useCase.Queue[:useCase.CountOfGoingRoutine], useCase.Queue[useCase.CountOfGoingRoutine:])
	useCase.Queue = useCase.Queue[useCase.CountOfGoingRoutine:]
	mutex.Unlock()
}

func (useCase *UseCase) onGoingOperation(body *models.Task, wg *sync.WaitGroup) {
	body.Status = "On going"
	body.StartTime = time.Now().Format("15:04:05 02.01.2006")
	currentValue := body.StartElement
	for i := 0; i < body.ElementCount; i++ {
		currentValue += body.Delta
		body.Iteration = i
		time.Sleep(time.Duration(int(body.L*1000)) * time.Millisecond)
	}
	body.Status = "Done"
	body.EndTime = time.Now().Format("15:04:05 02.01.2006")
	wg.Done()
}
