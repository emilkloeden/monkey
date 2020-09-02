let map = fn(arr, f) {
    let iter = fn(arr, acc) {
        if(len(arr) == 0) {
            acc
        } else {
            iter(rest(arr), push(acc, f(first(arr))))
        }
    };

    iter(arr, []);
};

let reduce = fn(arr, initial, f) {
    let iter = fn(arr, result) {
        if (len(arr) == 0) {
            result
        } else {
            iter(rest(arr), f(result, first(arr)));
        }
    };
    iter(arr, initial);
};



let sum = fn(arr) {
    reduce(arr, 0, fn(initial, el) {initial + el});
};

let arr = [ 1, 2, 3]
let doubled = map(arr, fn(x) {x * 2})
let summed = sum(doubled)

let m = 2 % 2 == 0

// a comment

let combined = arr + doubled;

let hash = {1: "hello", "two": "world"};
let hashKeys = keys(hash)
let hashValues = values(hash)

puts("Array: ")
puts(arr)
puts("Doubled: ")
puts(doubled)
puts("Summed: ")
puts(summed)
puts("Less than 5: ")
puts("TBD") // a comment afterwards
puts(m)
// final comment; puts(m);
puts("Combined: ")
puts(combined)

puts("Join 1 array:")
puts(join(arr, "~"))

puts("Join 2 arrays:")
puts(join([arr, doubled],"~"))

puts("Hash:")
puts(hash)

puts("Keys:")
puts(hashKeys)

puts("Values:")
puts(hashValues)