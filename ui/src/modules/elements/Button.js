import './Button.css'

const Button = ({textToDisplay, onClick}) => {
    return (
        <div className="button-container">
            <button className="button type1" onClick={onClick}>
                {textToDisplay}
            </button>
        </div>
    )

}

export default Button;