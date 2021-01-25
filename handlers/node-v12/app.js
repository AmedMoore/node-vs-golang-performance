const fs = require("fs");

function merge(left, right) {
    const array = [];
    while (left.length && right.length) {
        if (left[0] < right[0]) {
            array.push(left.shift());
        } else {
            array.push(right.shift());
        }
    }
    return [...array, ...left, ...right]
}

function mergeSort(array) {
    const half = array.length / 2;

    if (array.length < 2) {
        return array;
    }

    const left = array.splice(0, half);
    return merge(mergeSort(left), mergeSort(array));
}

exports.handler = async () => {
    const response = {};

    try {
        const data = fs.readFileSync("/opt/array.json");
        const array = JSON.parse(data.toString());

        const sortedArray = mergeSort(array);

        const res = [
            ...sortedArray.slice(0, 5),
            ...sortedArray.slice(sortedArray.length - 5, sortedArray.length)
        ];

        response.statusCode = 200;
        response.body = JSON.stringify(res);
    } catch (err) {
        response.statusCode = 500;
        response.body = JSON.stringify({error: err.message});
    }

    return response;
};
