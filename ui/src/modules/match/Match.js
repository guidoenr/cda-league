import React, { useState, useEffect } from 'react';
import Container from 'react-bootstrap/Container';
import PlayerCard from "../player/PlayerCard";
import Team from './Team'
import jsonPlayers from '../../resources/players.json'
import Button from '../elements/Button'
import './Match.css'
import 'animate.css';

const Match = () => {
    const [allPlayers, setAllPlayers] = useState([])
    let [availablePlayers, setAvailablePlayers] = useState([]);
    const [showTeams, setShowTeams] = useState(false);
    const [Team1, setTeam1] = useState({});
    const [Team2, setTeam2] = useState({});

    // obtain the players
    useEffect(() => {
        const getPlayers = async () => {
            try {
                const response = await fetch('http://localhost:8080/players/');
                const data = await response.json();
                setAllPlayers(data.players);
            } catch (error) {
                console.error(error);
            }
        };
        getPlayers();
    }, []);

    if (allPlayers.length === 0){
        setAllPlayers(jsonPlayers.players)
    }

    // generate match with selected players
    const generateMatchWithPlayers = () => {
        const data = { players: availablePlayers }
        fetch('http://localhost:8080/generateMatch/', {
            method: 'POST',
            body: JSON.stringify(data),
            headers: { 'Content-Type': 'application/json' }
        })
            .then(res => res.json())
            .then(response => {
                setTeam1(response.team1);
                setTeam2(response.team2);
                setShowTeams(true);
            })
            .catch(error => console.error('Error:', error));
    }


    // click the playerCard
    const handleSelectPlayer = (player) => {
        if (player.selected){
            availablePlayers.pop(player)
            player.selected=false
        } else {
            player.selected = true
            availablePlayers.push(player)
        }
    }


    const cdaLogo = require('../../assets/cda-league-only-logo.png')
    function renderTeams(){
        if (showTeams){
            return (
                <Container>
                <div className="match-container animate__animated animate__fadeInDown ">
                        <div className="match">
                            <div className="match-header">
                                <img className="match-logo" src={cdaLogo} alt="Carmen League"/>Carmen League<br/>
                                <div className="match-sub-header">
                                    Torneo de verano 2023
                                </div>
                            </div>

                            <div className="match-tournament"> </div>
                            <div className="match-content ">
                                <div className="team-container">
                                    <Team
                                        name={Team1.name}
                                        players={Team1.players}
                                        totalPlayers={Team1.totalPlayers}
                                        chanceOfWinning={Team1.chanceOfWinning}
                                    />
                                </div>
                                <div className="team-container">
                                    <Team
                                        name={Team2.name}
                                        players={Team2.players}
                                        totalPlayers={Team2.totalPlayers}
                                        chanceOfWinning={Team2.chanceOfWinning}
                                    />
                                </div>
                            </div>
                        </div>
                </div>
                </Container>
            )
        } else {
            return <Container></Container>
        }
    }

    return (
        <Container>
            <div className="note-container">
                <p className="match-note">Seleccionar jugadores</p>
            </div>
                <div className="available-players-container">
                    {allPlayers.map(player => (
                        <div
                            key={player.ID}
                            onClick={() => handleSelectPlayer(player)}>
                            <PlayerCard player={player} match={false}/>
                        </div>
                    ))}
                </div>
            <Container>
            <div className="button-container">
                <Button onClick={generateMatchWithPlayers} textToDisplay="Armar partido"/>
            </div>
            </Container>
            {renderTeams()}
        </Container>
    );
}
export default Match;
