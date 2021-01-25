package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"io/ioutil"
	"os"
)

func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}

func mergeSort(items []int) []int {
	var num = len(items)

	if num == 1 {
		return items
	}

	middle := num / 2
	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

func HandleRequest() (events.APIGatewayProxyResponse, error) {
	file, err := os.Open("/opt/array.json")
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       fmt.Sprintf("{ error: \"%s\" }", err.Error()),
		}, nil
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	byteValue, _ := ioutil.ReadAll(file)

	var array []int
	if err := json.Unmarshal(byteValue, &array); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       fmt.Sprintf("{ error: \"%s\" }", err.Error()),
		}, nil
	}

	sortedArray := mergeSort(array)

	var res []int
	res = append(res, sortedArray[0:5]...)
	res = append(res, sortedArray[len(sortedArray)-5:]...)
	if byteValue, err = json.Marshal(&res); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       fmt.Sprintf("{ error: \"%s\" }", err.Error()),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(byteValue),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
