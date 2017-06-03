var sum = 0;
var t1 = new Date();

for (var e=0; e<30; e++) {
    var x = [];
    for (var i=0; i<1000000; i++) {
        x.push(i);
    }

    var y = [];
    for (var i=0; i<1000000-1; i++) {
        y.push(x[i] + x[i+1]);
    }

    sum = 0;
    for (var i=0; i<1000000; i+=100) {
        sum += y[i];
    }
}

var t2 = new Date();
console.log('Elapsed (sec) ' + (t2-t1)/1000);
console.log(sum);
