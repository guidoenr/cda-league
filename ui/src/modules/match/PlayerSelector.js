import React, { useState, useEffect } from 'react';
import './PlayerSelector.css'
import PlayerCard from "../player/PlayerCard";
import Container from 'react-bootstrap/Container';

const PlayerSelector = ({ onTeamGeneration }) => {
    const [players, setPlayers] = useState([]);
    const [selectedPlayers, setSelectedPlayers] = useState([]);

    useEffect(() => {
        const fetchPlayers = async () => {
            try {
                const response = await fetch('http://localhost:8080/players/');
                const data = await response.json();
                setPlayers(data.players);
            } catch (error) {
                console.error(error);
            }
        };
        fetchPlayers();
    }, []);

    const handleSelectPlayer = (player) => {
        if (selectedPlayers.find(p => p.ID === player.ID)) {
            setSelectedPlayers(selectedPlayers.filter(p => p.ID !== player.ID))
        } else {
            setSelectedPlayers([...selectedPlayers, player])
        }
    }
    return (
        <Container>
        <h3>Available players</h3>
        <div className="available-players-container">
            {players.map(player => (
                <div
                    key={player.ID}
                    onClick={() => handleSelectPlayer(player)}>
                    <PlayerCard player={player}/>
                </div>
            ))}
        </div>
            <Container>
                <button className="btn-border btn btn4" onClick={() => onTeamGeneration(selectedPlayers)}>Armar Equipos</button>
            </Container>
        </Container>
    )
};

export default PlayerSelector;