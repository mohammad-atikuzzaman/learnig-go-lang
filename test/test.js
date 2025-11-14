const start = performance.now();
let sum = 0;
for (let i = 0; i <= 10000000; i++) {
  sum += i;
}
const end = performance.now();
console.log("Time :", end - start + " ms");
