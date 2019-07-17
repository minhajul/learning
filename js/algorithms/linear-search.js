
function linearSearch(array, number) {
    for (let i = 0; i < array.length; i++){
        if (array[i] === number){
            return i;
        }
    }

    return -1;
}

const array = [10, 4, 5, 1, 9, 7, 6];

console.log(linearSearch(array, 5));