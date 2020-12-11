const drink = {
  color: 'brown',
  carbonated: true,
  sugar: 40
};

const pepsi= ['brown',true,40];//thsi is array

const cocaCola: [string, boolean, number]= ['brown',true,40];//thsi is tupple

type Drink = [string,boolean, number];

const sprite: Drink=['brown',true,40]; // with new type


