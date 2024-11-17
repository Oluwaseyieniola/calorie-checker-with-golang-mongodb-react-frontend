import React, {useState, useEffect} from  'react';

import axios from "axios"; 

import {Button, Form, Container, Modal} from 'react-bootstrap'

import  Entry from './single-entry.components';

const Entries =() => {


    // my constants for the functions defined below in the code serving purposes like refresh, change entry, add new entry etc
    const [entries, setEntries] = useState([])
    const [refreshData, setRefreshData] = useState(false)
    const [changeEntry, setChangeEntry] = useState({'change': false, "id": 0})
    const [changeIngredient, setChangeIngredient]= useState({"change": false, "id": 0})
    const [newIngredientName, setNewIngredientName] = useState("")
    const [addnewEntry, setAddNewEntry] = useState(false)
    const [newEntry, setNewEntry] =useState({"dish": "", "ingredients": "", "calories":0, fat:0})


    // the initial load up
    useEffect(() => {
        getAllEntries();
    }, [])

    if(refreshData){
        setRefreshData(false);
        getAllEntries();
    }
    return (
        <div>
            <Container> 
<Button>onClick ={()  => setAddNewEntry(true)}Track your day's calories</Button>
            </Container>
            <Container>
{entries != null && entries.map((entry, i) =>(
    <Entry entryData= {entry} deleteSingleEntry={deleteSingleEntry} setChangeIngredient={setChangeIngredient} setChangeEntry= {setChangeEntry}/>
))}
            </Container>
            <Modal show ={addnewEntry} onHide={() => setAddNewEntry(false)} centered>
                <Modal.Header closeButton>
<Modal.Title Add Calorie Entry>

</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <Form.Group>
                        <Form.Label>
                            dish
                        </Form.Label>
                        <Form.Control onChange={(event) => {newEntry.dish = event.target.value}}></Form.Control>
                        <Form.Label>
                            ingredients
                        </Form.Label>
                        <Form.Control onChange={(event) => {newEntry.ingredients = event.target.value}}></Form.Control>
                        <Form.Label>
                            calories
                        </Form.Label>
                        <Form.Control onChange={(event) => {newEntry.calories = event.target.value}}></Form.Control>
                        <Form.Label>
                            fat
                        </Form.Label>
                        <Form.Control type = "number" onChange={(event) => {newEntry.fat = event.target.value}}></Form.Control>
                    </Form.Group>
                    <Button  onClick={() => addSingNewEntry()}>Add</Button>
                    <Button onClick={() =>  setAddNewEntry(false)}>Cancel </Button>
                </Modal.Body>

            </Modal>
            <Modal show ={changeIngredient.change} onHide= {() =>  setChangeIngredient({"change": false, "id": 0})} centered></Modal>
            <Modal.Header closeButton>
                <Modal.Title> Change ingredients</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <Form.Group>
                    <Form.Label> new ingredients</Form.Label>
                    <Form.Control> onChange={event => {setNewIngredientName(event.target.value)}}</Form.Control>
                </Form.Group>
                <Button onClick= {() => changeIngredientForEntry()}>Change </Button>
                <Button onClick= {() => changeIngredientForEntry({"change": false, "id":0})}>Cancel </Button>
            </Modal.Body>
        

        <Modal show = {changeEntry.change } onHide={() => setChangeEntry({"change": false, "id":0})} centered>
            <Modal.Header closeButton>
                <Modal.Title> Change Entry</Modal.Title>
                 </Modal.Header>
                 <Modal.Body>
                    <Form.Group>
                        <Form.Label>dish</Form.Label>
                        <Form.Control onChange={(event) => {newEntry.dish = event.target.value}}> </Form.Control>
                        <Form.Label>ingredients</Form.Label>
                        <Form.Control onChange={(event) => {newEntry.ingredients = event.target.value}}> </Form.Control>

                        <Form.Label>calories</Form.Label>
                        <Form.Control onChange={(event) => {newEntry.calories = event.target.value}}> </Form.Control>

                        <Form.Label>fat</Form.Label>
                        <Form.Control type = "number" onChange={(event) => {newEntry.fat = event.target.value}}> </Form.Control>
                    </Form.Group>
                    <Button onClick={() => changeSingleEntry()}>change</Button>
                    <Button onClick={() => setChangeEntry({"change": false, "id": 0})}>Cancel </Button>
                 </Modal.Body>
        
</Modal>
        </div>
    );
    function changeIngredientForEntry(){
        changeIngredient.change = false;
        var url = "http://localhost:3030/ingredient/update/"  + changeIngredient.id
        axios.put(url, {
            "ingredients": newIngredientName
        })
        .then(response => {
            console.log(response.status)
            if (response.status == 200){
                setRefreshData(true)
            }
        })
    }

    function changeSingleEntry(){
        changeEntry.change = false;
        var url = "http://localhost:3030/entry/update/" + changeEntry.id
        axios.put(url, newEntry)
        .then(response => {
            if (response.status == 200){
                setRefreshData(true)
            }
        })
    }
    function addSingNewEntry(){
        setAddNewEntry(false) 
        var url = "http://localhost:3030/entry/create/"
        axios.post(url, {
            "ingredients": newEntry.ingredients,
            "dish": newEntry.dish,
            "calories": newEntry.calories,
            "fat":parseFloat(newEntry.fat),

        }).then(response => {
            if(response.status == 200){
                setRefreshData(true)
            }
        })
    }
}
function deleteSingleEntry(id){
    var url = "http://localhost:3030/entry/delete/" +id
    axios.delete(url).then(response => {
        if (response.status == 200){
            setRefreshData(true)
        }
    })
}

function getAllEntries(){
    var url = "http://localhost:3030/entries/"
    axios.get(url, {
        responseType: 'json'

    }).then(response => {
        if (response.status == 200){
            setEntries(response.data)
        }
    })
}


export default Entries
    