import React, {
    useState
} from 'react';
import { Button, Input } from 'antd';
const list = [{
        id: 0,
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
    const [search, setSearch] = useState('')
    console.log(search)
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
            Button key = {-1 }
            onClick = { handleAdd }
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
        Input value = {
            search
        }
        onChange = {
            e => setSearch(e.target.value)
        }
        placeholder = "Search..." / >
        <
        div > {
            renderList()

        } < /div>

        <
        /div >

    )

}

export default App;