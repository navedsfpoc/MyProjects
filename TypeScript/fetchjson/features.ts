const test: string = 'test';

//when to use annotations:
//1) Function that returns the any  type

const json = '{"x":10,"y":20}';
const coordinates = JSON.parse(json);
console.log(coordinates);

let testOne: string;

let words = ['red','green'];
let foundWord: boolean | number = false;

for(let i =0; i< words.length;i++){

}

const add = (a: number, b:number):  number =>{
  return a+b;
}

console.log(add(2,3));

function divide(){

}

