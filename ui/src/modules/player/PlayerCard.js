import React, {useState} from 'react';
import Container from 'react-bootstrap/Container';
import {getStars, getPhoto, getColor} from "../Util";
import './PlayerCard.css'
import '../Util'


const PlayerCard = ({player, match} ) => {
    const [isSelected, setIsSelected] = useState(false);
    const handleClick = () => {
        setIsSelected(!isSelected);
    }

    const selected = isSelected ? " selected" : "";

    function snakeBorder(){
        if (selected){
            return (
                <div>
                    <span/>
                    <span/>
                    <span/>
                    <span/>
                </div>
            )
        }
    }


    function render(){
        // if the player will be displayed in their match form
        if (match){
            return (
                <Container className={"player-card-match"}>
                    <img src={getPhoto(player.nickname)} alt={`${player.name}'s profile photo`} className="photo" />
                    <div className="info">
                        <h2 className="nickname">{player.nickname}</h2>
                        <div className="rank">
                            {getStars(player.rank)}
                        </div>
                        <div className="position-shape" style={{backgroundColor: getColor(player.position)}}> {player.position.toUpperCase().substring(0,3)}</div>
                    </div>
                </Container>
            );
        } else {
            return (
                <Container className={"player-card" + selected} onClick={handleClick} data-position={player.position} >
                    {snakeBorder()}
                    <img src={getPhoto(player.nickname)} alt={`${player.name}'s profile photo`} className="photo" />
                    <div className="info">
                        <h2 className="nickname">{player.nickname}</h2>
                        <div className="rank">
                            {getStars(player.rank)}
                        </div>
                        <div className="name">{player.name}</div>
                        <div className="elo">ELO: {player.elo}</div>
                        <div className="position-shape" style={{backgroundColor: getColor(player.position)}}> {player.position.toUpperCase().substring(0,3)}</div>
                    </div>
                </Container>
            );
        }
    }

    return render()
}

export default PlayerCard;