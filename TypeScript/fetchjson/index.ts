import axios from 'axios';

const url = 'https://jsonplaceholder.typicode.com/todos/1';

interface Todo {
  id: number;
  title: string;
  completed: boolean;
}
axios.get(url).then(response=>{
  console.log(response.data.id);
  const todo = response.data as Todo;
  const ID = todo.id;
  const title = todo.title;
  const finished = todo.completed;
  logTodo (ID, title);
});

const logTodo =(id, title)=>{
  console.log(`
  ${id}
  ${title}
  `);

};
