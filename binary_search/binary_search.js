function binarySearch(arr, number) {
    let low = 0
    let high = arr.length - 1
    let count = 0

    while(low <= high) {
        count++
        let middle = Math.floor((low + high) / 2)
        let guess = arr[middle]

        if(guess === number) {
            console.log(`Performed ${count} actions`)
            return middle
        }
        if(guess < number) {
            low = middle + 1
        }
        if(guess > number) {
            high = middle - 1
        }
    }
    console.log(`Performed ${count} actions`)
    return 
}

const arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
console.log(binarySearch(arr, 7))
