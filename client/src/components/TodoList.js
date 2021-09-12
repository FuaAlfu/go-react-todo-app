import React, {Component} from 'react';
import axios from 'axios';
import {Card, Header, Form, Input, Icon} from 'semantic-ui-react';

let endpoint = "http://localhost:9090";

class TodoList extends Component{
    constructor(props){
        super(props);

        this.state = {
            task: "",
            items: [],
        };
    }
    ComponentDidMount(){
        this.getTask();
    }

    render(){
        return(
            <>
              <div className="row">
                <Header className="header" as="h2" color="red">
                   Todo List
                </Header>
              </div>
              <div className="row">
                 <Form onSubmit={this.onSubmit}>
                     <Input
                     type="text"
                     name="task"
                     onChange={this.onChange}
                     value={this.state.task}
                     fluid
                     placeholder="Create Task"
                     />
                 </Form>
              </div>
            </>
        )
    }
}

export default TodoList;