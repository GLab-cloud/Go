import React from 'react';
import { Button, Input } from 'antd';
const list = [{
        id: 1,
        title: 'Code'
    },
    {
        id: 1,
        title: 'Eat'
    }, {
        id: 2,
        title: 'Sleep'
    }
]

function App() {
    const renderList = () => {
        const l = list.map(item => {
            return ( <
                div key = {
                    item.id
                } > {
                    item.title
                } < /div>

            )
        })
        l.push( <
            Button onClick = { handleAdd }
            type = "primary" > Add job < /Button>
        )
        return l
    }
    const handleAdd = () => {
        console.log('handle Add open dialog')
    }
    return ( <
        div className = 'container' >
        <
        Input placeholder = "please enter something..." / >
        <
        div > {
            renderList()

        } < /div>

        <
        /div >

    )

}

export default App;