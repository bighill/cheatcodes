const thePromise = (message, delayInMS) =>
  new Promise((resolve) => {
    setTimeout(() => resolve(message), delayInMS);
  });

const doThePromise = async () => {
  const a = await thePromise("Waiting...", 333);
  console.log(a);
  return await thePromise("Done.", 555);
};

doThePromise().then((res) => {
  console.log(res);
});

console.log("Starting...");
