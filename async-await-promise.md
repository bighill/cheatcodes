# Using async/await with a promise

This bit of code should explain it.

```js
function doubleAfter2Seconds(x) {
  return new Promise(resolve => {
    setTimeout(() => {
      resolve(x * 2);
    }, 2000);
  });
}

async function addAsync(x) {
  const a = await doubleAfter2Seconds(10);
  console.log(a)
  const b = await doubleAfter2Seconds(20);
  console.log(b)
  const c = await doubleAfter2Seconds(30);
  console.log(c)
  return x + a + b + c;
}


addAsync(10).then((sum) => {
  console.log(sum);
});
```
