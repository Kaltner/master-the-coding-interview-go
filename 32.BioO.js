// What is the Big O of the below function? (Hint, you may want to go line by line)
// Answer: O(N)
// Actual Answer: O(3 + 4n)
function funChallenge(input) {
  let a = 10;
  a = 50 + 3;

  for (let i = 0; i < input.length; i++) {
    anotherFunction();
    let stranger = true;
    a++;
  }
  return a;
}
console.log(funChallenge(50))