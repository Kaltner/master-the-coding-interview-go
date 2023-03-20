// Log all pair of arrays
const boxes = [1,2,3,4,5]
boxes.forEach(e => {
    boxes.forEach(e2 => {
        if( e === e2) {
            console.log(`Found pair: ${e} ${e2}`)
        }
    })
});