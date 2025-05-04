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
        return l
    }
    const handleAdd = () => {
        console.log('handleAdd')
    }
    return (
        // <Button type="primary" > Primary Button < /Button>
        <
        div className = 'container' >
        <
        Input placeholder = "please enter something..." / >
        <
        div > {
            renderList()

        } < /div>

        <
        Button onClick = { handleAdd }
        type = "primary" > Add job < /Button> < /div >

    )

}

export default App;