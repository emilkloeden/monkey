let arrays = import("../stdlib/arrays.mo")

let myArr = [1,2,3,4,5];

puts(arrays["sum"](myArr))

let isEven = fn(x) { x % 2 == 0 }
puts(keys(arrays))
puts(arrays["filter"](myArr, isEven))