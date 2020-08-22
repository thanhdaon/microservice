
function add(a, b) {
  if (b) {
    return a + b;
  }

  return (n) => a + n
}


console.log(add(1, 2))
console.log(add(1)(2))