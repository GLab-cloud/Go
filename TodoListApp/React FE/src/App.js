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
        return list.map(item => {
            return ( <
                div key = {
                    item.id
                } > {
                    item.title
                } < /div>

            )
        })
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
        Button type = "primary" > Add job < /Button> < /div >

    )

}

export default App;