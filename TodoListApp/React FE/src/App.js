import React, {
    useState
} from 'react';
import { Input, Modal, Checkbox } from 'antd';
import {
    debounce
}
from 'lodash';
const initList = [{
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
    const [job, setJob] = useState("")
    const [list, setList] = useState(initList)

    const renderList = () => {
        const l = list.map(item => {
            return ( <
                div key = {
                    item.id
                }
                className = 'mt-2 border p-2' > < div > {
                    item.title
                } < /div> <
                div className = 'mt-1' > < Checkbox onChange = { onChange } >
                Done < /Checkbox > < /div > < /div >

            )
        })
        l.push( <
            // Button key = {-1 }
            // onClick = { handleAdd }
            // type = "primary"
            // className = 'border bg-info text-white d-flex justify-content-center' >
            // Add job < /Button>
            div key = {-1 }
            onClick = { handleAdd }

            className = 'mt-2 border bg-info text-white d-flex justify-content-center' >
            Add Job < /div>
        )
        return l
    }
    const onChange = () => {}

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
        if (job.length > 0) {

            const item = {
                id: list.length,
                title: job
            }
            setList([item, ...list])
            setVisible(false)
            setJob('')
                // call API set job to server Golang...
        }

    }
    const handleCreateModalCancel = () => {
        setVisible(false)
        setJob('')
    }
    return ( < > <
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

        } < /div> < /
        div > < Modal title = "Create A New Job"
        visible = { visible }
        onOk = {
            handleCreateModalOk
        }
        onCancel = {
            handleCreateModalCancel
        } >
        <
        div >
        <
        label > What do you do ? < /label> <
        Input
        value = { job }
        onChange = { e => setJob(e.target.value) }
        placeholder = "Enter something..." / >
        <
        /div> < /
        Modal >
        <
        />

    )

}

export default App;