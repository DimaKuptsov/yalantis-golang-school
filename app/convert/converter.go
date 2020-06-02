package converter

import (
	"errors"
	"fmt"
)

func ConvertMultidimensionalSlice(sliceForConvert [][]int) ([]int, error) {
	var convertedSlice []int
	var error error
	sliceForConvertLength := len(sliceForConvert)
	for _, internalSlice := range sliceForConvert {
		if len(internalSlice) != sliceForConvertLength {
			errorMessage := fmt.Sprintf("Internal slices must have len %v. Internal slice with bad len: %v", sliceForConvertLength, internalSlice)
			error = errors.New(errorMessage)
		}
	}
	if error == nil {
		var sliceCircleCut []int
		for canCutSliceByCircle := true; canCutSliceByCircle; canCutSliceByCircle = len(sliceForConvert) > 0 {
			sliceForConvert, sliceCircleCut = cutSliceByCircle(sliceForConvert)
			convertedSlice = append(convertedSlice, sliceCircleCut...)
		}
	}
	return convertedSlice, error
}

func cutSliceByCircle(sliceForCut [][]int) ([][]int, []int) {
	var cuttedSlice []int
	cuttedSlice = append(sliceForCut[0])
	sliceForCut = sliceForCut[1:]
	sliceForCutLength := len(sliceForCut)
	if sliceForCutLength > 0 {
		for internalSliceIndex := 0; internalSliceIndex < sliceForCutLength-1; internalSliceIndex++ {
			internalSlice := sliceForCut[internalSliceIndex]
			internalSliceLength := len(internalSlice)
			if internalSliceLength > 0 {
				lastSliceElementIndex := internalSliceLength - 1
				cuttedSlice = append(cuttedSlice, internalSlice[lastSliceElementIndex])
				sliceForCut[internalSliceIndex] = sliceForCut[internalSliceIndex][:lastSliceElementIndex]
			}
		}
		lastInternalSliceIndex := sliceForCutLength - 1
		revertedLastInternalSlice := revertSlice(sliceForCut[lastInternalSliceIndex])
		cuttedSlice = append(cuttedSlice, revertedLastInternalSlice...)
		sliceForCut = sliceForCut[:lastInternalSliceIndex]
		sliceForCutLength = len(sliceForCut)
		for internalSliceIndex := sliceForCutLength - 1; internalSliceIndex >= 1; internalSliceIndex-- {
			internalSlice := sliceForCut[internalSliceIndex]
			internalSliceLength := len(internalSlice)
			if internalSliceLength > 0 {
				cuttedSlice = append(cuttedSlice, internalSlice[0])
				sliceForCut[internalSliceIndex] = sliceForCut[internalSliceIndex][1:]
			}
		}
	}
	return sliceForCut, cuttedSlice
}

func revertSlice(sliceForRevert []int) []int {
	for elementIndex := len(sliceForRevert)/2 - 1; elementIndex >= 0; elementIndex-- {
		oppositeElementIndex := len(sliceForRevert) - 1 - elementIndex
		sliceForRevert[elementIndex], sliceForRevert[oppositeElementIndex] = sliceForRevert[oppositeElementIndex], sliceForRevert[elementIndex]
	}
	return sliceForRevert
}
