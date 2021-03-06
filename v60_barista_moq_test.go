// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package v60

import (
	"sync"
)

// Ensure, that V60BaristaMock does implement V60Barista.
// If this is not the case, regenerate this file with moq.
var _ V60Barista = &V60BaristaMock{}

// V60BaristaMock is a mock implementation of V60Barista.
//
//     func TestSomethingThatUsesV60Barista(t *testing.T) {
//
//         // make and configure a mocked V60Barista
//         mockedV60Barista := &V60BaristaMock{
//             AddWaterFunc: func(amountOfWater float64)  {
// 	               panic("mock out the AddWater method")
//             },
//             BloomCoffeeFunc: func(amountOfWater float64)  {
// 	               panic("mock out the BloomCoffee method")
//             },
//             PrepV60Func: func()  {
// 	               panic("mock out the PrepV60 method")
//             },
//             StirFunc: func()  {
// 	               panic("mock out the Stir method")
//             },
//         }
//
//         // use mockedV60Barista in code that requires V60Barista
//         // and then make assertions.
//
//     }
type V60BaristaMock struct {
	// AddWaterFunc mocks the AddWater method.
	AddWaterFunc func(amountOfWater float64)

	// BloomCoffeeFunc mocks the BloomCoffee method.
	BloomCoffeeFunc func(amountOfWater float64)

	// PrepV60Func mocks the PrepV60 method.
	PrepV60Func func()

	// StirFunc mocks the Stir method.
	StirFunc func()

	// calls tracks calls to the methods.
	calls struct {
		// AddWater holds details about calls to the AddWater method.
		AddWater []struct {
			// AmountOfWater is the amountOfWater argument value.
			AmountOfWater float64
		}
		// BloomCoffee holds details about calls to the BloomCoffee method.
		BloomCoffee []struct {
			// AmountOfWater is the amountOfWater argument value.
			AmountOfWater float64
		}
		// PrepV60 holds details about calls to the PrepV60 method.
		PrepV60 []struct {
		}
		// Stir holds details about calls to the Stir method.
		Stir []struct {
		}
	}
	lockAddWater    sync.RWMutex
	lockBloomCoffee sync.RWMutex
	lockPrepV60     sync.RWMutex
	lockStir        sync.RWMutex
}

// AddWater calls AddWaterFunc.
func (mock *V60BaristaMock) AddWater(amountOfWater float64) {
	if mock.AddWaterFunc == nil {
		panic("V60BaristaMock.AddWaterFunc: method is nil but V60Barista.AddWater was just called")
	}
	callInfo := struct {
		AmountOfWater float64
	}{
		AmountOfWater: amountOfWater,
	}
	mock.lockAddWater.Lock()
	mock.calls.AddWater = append(mock.calls.AddWater, callInfo)
	mock.lockAddWater.Unlock()
	mock.AddWaterFunc(amountOfWater)
}

// AddWaterCalls gets all the calls that were made to AddWater.
// Check the length with:
//     len(mockedV60Barista.AddWaterCalls())
func (mock *V60BaristaMock) AddWaterCalls() []struct {
	AmountOfWater float64
} {
	var calls []struct {
		AmountOfWater float64
	}
	mock.lockAddWater.RLock()
	calls = mock.calls.AddWater
	mock.lockAddWater.RUnlock()
	return calls
}

// BloomCoffee calls BloomCoffeeFunc.
func (mock *V60BaristaMock) BloomCoffee(amountOfWater float64) {
	if mock.BloomCoffeeFunc == nil {
		panic("V60BaristaMock.BloomCoffeeFunc: method is nil but V60Barista.BloomCoffee was just called")
	}
	callInfo := struct {
		AmountOfWater float64
	}{
		AmountOfWater: amountOfWater,
	}
	mock.lockBloomCoffee.Lock()
	mock.calls.BloomCoffee = append(mock.calls.BloomCoffee, callInfo)
	mock.lockBloomCoffee.Unlock()
	mock.BloomCoffeeFunc(amountOfWater)
}

// BloomCoffeeCalls gets all the calls that were made to BloomCoffee.
// Check the length with:
//     len(mockedV60Barista.BloomCoffeeCalls())
func (mock *V60BaristaMock) BloomCoffeeCalls() []struct {
	AmountOfWater float64
} {
	var calls []struct {
		AmountOfWater float64
	}
	mock.lockBloomCoffee.RLock()
	calls = mock.calls.BloomCoffee
	mock.lockBloomCoffee.RUnlock()
	return calls
}

// PrepV60 calls PrepV60Func.
func (mock *V60BaristaMock) PrepV60() {
	if mock.PrepV60Func == nil {
		panic("V60BaristaMock.PrepV60Func: method is nil but V60Barista.PrepV60 was just called")
	}
	callInfo := struct {
	}{}
	mock.lockPrepV60.Lock()
	mock.calls.PrepV60 = append(mock.calls.PrepV60, callInfo)
	mock.lockPrepV60.Unlock()
	mock.PrepV60Func()
}

// PrepV60Calls gets all the calls that were made to PrepV60.
// Check the length with:
//     len(mockedV60Barista.PrepV60Calls())
func (mock *V60BaristaMock) PrepV60Calls() []struct {
} {
	var calls []struct {
	}
	mock.lockPrepV60.RLock()
	calls = mock.calls.PrepV60
	mock.lockPrepV60.RUnlock()
	return calls
}

// Stir calls StirFunc.
func (mock *V60BaristaMock) Stir() {
	if mock.StirFunc == nil {
		panic("V60BaristaMock.StirFunc: method is nil but V60Barista.Stir was just called")
	}
	callInfo := struct {
	}{}
	mock.lockStir.Lock()
	mock.calls.Stir = append(mock.calls.Stir, callInfo)
	mock.lockStir.Unlock()
	mock.StirFunc()
}

// StirCalls gets all the calls that were made to Stir.
// Check the length with:
//     len(mockedV60Barista.StirCalls())
func (mock *V60BaristaMock) StirCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockStir.RLock()
	calls = mock.calls.Stir
	mock.lockStir.RUnlock()
	return calls
}
