import React, { useState, useEffect } from 'react';
import './PlayerSelector.css'

const PlayerSelector = ({ onTeamGeneration }) => {
    const [players, setPlayers] = useState([]);
    const [selectedPlayers, setSelectedPlayers] = useState([]);

    useEffect(() => {
        const fetchPlayers = async () => {
            try {
                const response = await fetch('http://localhost:8080/players/');
                const data = await response.json();
                setPlayers(data);
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
        <div className="player-selector">
            <h3>Players Available</h3>
            <div className="players-container">
                {players.map(player => (
                    <div
                        key={player.ID}
                        className={`player ${selectedPlayers.find(p => p.ID === player.ID) && 'selected'}`}
                        onClick={() => handleSelectPlayer(player)}
                    >
                        <p>{player.name}</p>
                    </div>
                ))}
            </div>
            <button onClick={() => onTeamGeneration(selectedPlayers)}>Generate teams</button>
        </div>
    )
};

export default PlayerSelector;