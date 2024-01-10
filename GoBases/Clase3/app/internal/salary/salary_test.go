package salary_test

import (
	"app/internal/salary"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalcularImpuesto(t *testing.T) {
	t.Run("Success - case 01 - salary less than 50000", func(t *testing.T) {
		// arrange

		// act
		s := 50000.0
		result, err := salary.CalcularImpuesto(s)
		// assert
		expectedError := ""
		expectedResult := 0.0

		require.Equal(t, expectedError, err)
		require.Equal(t, expectedResult, result)

	})

	t.Run("Success - case 02 - salary greater than 50000", func(t *testing.T) {
		// arrange

		// act
		s := 100000.0
		result, err := salary.CalcularImpuesto(s)
		// assert
		expectedError := ""
		expectedResult := 17000.0

		require.Equal(t, expectedError, err)
		require.Equal(t, expectedResult, result)

	})

	t.Run("Success - case 03 - salary greater than 150000", func(t *testing.T) {
		// arrange

		// act
		s := 160000.0
		result, err := salary.CalcularImpuesto(s)
		// assert
		expectedError := ""
		expectedResult := 43200.0

		require.Equal(t, expectedError, err)
		require.Equal(t, expectedResult, result)

	})
}
