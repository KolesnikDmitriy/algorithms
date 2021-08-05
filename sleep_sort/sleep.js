const a = [4, 1, 3, 2]

a.forEach(el => {
    setTimeout(function() {
        console.log(el);
    }, el);
});