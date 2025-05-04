import React, {
    //useState
} from 'react';
import { Button, Input } from 'antd';
import {
    debounce
}
from 'lodash';
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
    //const [search] = useState('')
    //console.log(search)
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
    const handleSearchDebounced = debounce((q) => {
        //call API search...
        console.log('handleSearchDebounced', q)

    }, 500)
    const handleSearch = (e) => { handleSearchDebounced(e.target.value) }
    return ( <
        div className = 'container' >
        <
        // Input value = {
        //     search
        // }
        Input onChange = {
            // e => setSearch(e.target.value)
            handleSearch
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