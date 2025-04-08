import React from 'react';
import Confetti from 'react-confetti';

interface ConfettiComponentProps {
  width: number;
  height: number;
  runConfetti: boolean;
  gravity?: number;
  numberOfPieces?: number;
}

const ConfettiComponent: React.FC<ConfettiComponentProps> = ({
  width,
  height,
  runConfetti,
  gravity = 1.0,  
  numberOfPieces = 175,
}) => {
  return (
    <Confetti
      width={width}
      height={height}
      run={runConfetti}  
      gravity={gravity}  
      numberOfPieces={numberOfPieces} 
      recycle={false}  
      initialVelocityX={{ min: -10, max: 10 }} 
      initialVelocityY={{ min: 30, max: 50 }} 
    />
  );
};

export default ConfettiComponent;
