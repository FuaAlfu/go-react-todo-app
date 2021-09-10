import React from 'react';
import './App.css';
import {Container} from 'semantic-ui-react';
import TodoList from './components/TodoList'

function App() {
  return (
    <div className="App">
      <Container>
         <TodoList />
      </Container>
    </div>
  );
}

export default App;
