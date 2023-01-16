import React, { useState, useEffect } from 'react';
import Team from './Team'
import Container from 'react-bootstrap/Container';
import './Match.css'
import Button from "@mui/material/Button";
import PlayerCard from "../player/PlayerCard";

const Match = () => {
    const [showTeams, setShowTeams] = useState(false);
    const [availablePlayers, setAvailablePlayers] = useState([]);
    const [Team1, setTeam1] = useState({});
    const [Team2, setTeam2] = useState({});

    // obtain the players
    useEffect(() => {
        const getPlayers = async () => {
            try {
                const response = await fetch('http://localhost:8080/players/');
                const data = await response.json();
                setAvailablePlayers(data.players);
            } catch (error) {
                console.error(error);
            }
        };
        getPlayers();
    }, []);

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
        if (availablePlayers.find(p => p.ID === player.ID)) {
            setAvailablePlayers(availablePlayers.filter(p => p.ID !== player.ID))
        } else {
            setAvailablePlayers([...availablePlayers, player])
        }
    }


        return (
            <Container>
                <Container>
                    <h3>Available players</h3>
                    <div className="available-players-container">
                        {availablePlayers.map(player => (
                            <div
                                key={player.ID}
                                onClick={() => handleSelectPlayer(player)}>
                                <PlayerCard player={player}/>
                            </div>
                        ))}
                    </div>
                </Container>
                <Container>
                    <Button onClick={() => generateMatchWithPlayers()} variant="contained" className="btn-color">Armar Match</Button>
                </Container>
                <div className="match-container">
                    <Team className="team-container"
                        name="TEAM 1"
                        players={Team1.players}
                        totalPlayers={Team1.totalPlayers}
                        chanceOfWinning={Team1.chanceOfWinning}
                    />
                    <Team className="team-container"
                        name="TEAM 2"
                        players={Team2.players}
                        totalPlayers={Team2.totalPlayers}
                        chanceOfWinning={Team2.chanceOfWinning}
                    />
                </div>
            </Container>
        );
}



export default Match;
