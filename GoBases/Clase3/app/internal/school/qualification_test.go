package school_test

import (
	"app/internal/school"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAverage(t *testing.T) {
	t.Run("Success - case 01 - average qualifications ", func(t *testing.T) {
		// arrange

		// act
		notas := []float64{1, 2, 3, 4, 5, 10, 10, 10, 10, 10}
		result, err := school.Average(notas...)
		// assert
		expectedError := ""
		expectedResult := 6.5

		require.Equal(t, expectedError, err)
		require.Equal(t, expectedResult, result)
	})

	t.Run("Failure - case 02 - empty qualifications ", func(t *testing.T) {
		// arrange

		// act
		notas := []float64{}
		result, err := school.Average(notas...)
		// assert
		expectedError := "Notas vacias"
		expectedResult := 0.0

		require.Equal(t, expectedError, err)
		require.Equal(t, expectedResult, result)
	})

}
