import React, {
    useState
} from 'react';
import { Button, Input, Modal } from 'antd';
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
    const [visible, setVisible] = useState(false)
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
        setVisible(true)
    }
    const handleSearchDebounced = debounce((q) => {
        //call API search...
        console.log('handleSearchDebounced', q)

    }, 500)
    const handleSearch = (e) => { handleSearchDebounced(e.target.value) }
    const handleCreateModalOk = () => {
        setVisible(false)
    }
    const handleCreateModalCancel = () => {
        setVisible(false)
    }
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

        } < /div> <
        Modal title = "Create"
        visible = { visible }
        onOk = {
            handleCreateModalOk
        }
        onCancel = {
            handleCreateModalCancel

        } > <
        p > Some contents... < /
        p >
        <
        p > Some contents... < /
        p >
        <
        p > Some contents... < /
        p >
        <
        /
        Modal >
        <
        /
        div >

    )

}

export default App;