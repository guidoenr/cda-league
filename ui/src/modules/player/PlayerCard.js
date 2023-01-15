import React from 'react';
import './PlayerCard.css'
import '../Util'
import {getStars, getPhoto, getColor} from "../Util";
import Container from 'react-bootstrap/Container';

const PlayerCard = ({player}) => {

    return (
        <Container className="player-card">
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

export default PlayerCard;