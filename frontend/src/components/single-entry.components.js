import React, {useState, useEffect} from  'react';

import 'bootstrap/dist/css/bootstrap.css'


import {Button, Card, Row, Col} from 'react-bootstrap'
// in this code each entry connotes dish, ingredients, calories, fat etc entered by the user

const Entry  = ({entryData, setChangeIngredient, deleteSingleEntry, setChangeEntry}) => {
    return (
        <Card>
            <Row>
                <Col>Dish: {entryData !== undefined && entryData.dish}</Col>
                <Col>Ingredients: {entryData !== undefined && entryData.ingredients}</Col>
                <Col>Calories: {entryData !== undefined && entryData.calories}</Col>
                <Col>Fat: {entryData !== undefined && entryData.fat}</Col>
                <Col><Button onClick={()=> deleteSingleEntry}> delete entry</Button></Col>
                <Col><Button onClick={()=> changeIngredient()}> change ingredients</Button></Col>
                <Col><Button onClick={()=> changeEntry()}> change entry</Button></Col>
            </Row>
        </Card>
    )


    // so here, users should be able to change ingredient in each entry as they wish to
    function changeIngredient(){
        setChangeIngredient (
            {
                "change": true, 
                "id": entryData._id
            }
        )
    }
    function changeEntry(){
        setChangeEntry(
            {
                "change": true, 
                "id": entryData._id
            }
        )
    }
    
}

export default  Entry