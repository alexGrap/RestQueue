package usecase

import (
	models "inter/internal"
	"sort"
	"sync"
	"time"
)

type UseCase struct {
	InMemory map[int]*models.Task
}

var CountOfGoingRoutine int

var queue []int
var lengthQueue int

func InitUseCase() models.UseCase {
	useCase := UseCase{}
	lengthQueue = 0
	useCase.InMemory = make(map[int]*models.Task)
	return &useCase
}

func (useCase *UseCase) Create(task models.Task) {
	var mutex sync.Mutex
	appender := func(position int, mutex *sync.Mutex) {
		mutex.Lock()
		queue = append(queue, position)
		mutex.Unlock()
	}
	task.Place = len(useCase.InMemory)
	lengthQueue += 1
	task.Status = "In queue"
	task.CreationTime = time.Now().Format("15:04:05 02.01.2006")
	useCase.InMemory[task.Place] = &task
	appender(task.Place, &mutex)
	if lengthQueue >= CountOfGoingRoutine {
		go useCase.Handle()
		lengthQueue -= CountOfGoingRoutine
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

func (useCase *UseCase) Handle() {
	var wg sync.WaitGroup
	sort.Slice(queue, func(i, j int) bool {
		return queue[i] < queue[j]
	})
	operationFunc := func(body *models.Task) {
		defer wg.Done()
		body.Status = "On going"
		body.StartTime = time.Now().Format("15:04:05 02.01.2006")
		currentValue := body.StartElement
		for i := 0; i < body.ElementCount; i++ {
			currentValue += body.Delta
			time.Sleep(time.Duration(int(body.L*1000)) * time.Millisecond)
		}
		body.Status = "Done"
		TTL := time.Duration(int(body.L*1000)) * time.Millisecond
		timer := time.AfterFunc(TTL, func() {
			delete(useCase.InMemory, body.Place)
		})
		defer timer.Stop()
	}
	wg.Add(CountOfGoingRoutine)
	for i := 0; i < CountOfGoingRoutine; i++ {
		go operationFunc(useCase.InMemory[queue[i]])

	}
	wg.Wait()
	deleter()
}

func deleter() {
	if len(queue) == CountOfGoingRoutine {
		queue = nil
		return
	}
	copy(queue[:CountOfGoingRoutine], queue[CountOfGoingRoutine:])
	queue = queue[:len(queue)-CountOfGoingRoutine]
}
