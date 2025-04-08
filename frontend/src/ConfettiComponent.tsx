// ConfettiComponent.tsx
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
  gravity = 1.0,  // Higher gravity to ensure the confetti falls quickly
  numberOfPieces = 175, // More pieces for a dramatic effect
}) => {
  return (
    <Confetti
      width={width}
      height={height}
      run={runConfetti}  // Starts when runConfetti is true
      gravity={gravity}  // Increased gravity for faster fall
      numberOfPieces={numberOfPieces}  // More confetti pieces
      recycle={false}  // Prevent confetti from recycling (no reappearing pieces)
      initialVelocityX={{ min: -10, max: 10 }} // Spread horizontally
      initialVelocityY={{ min: 30, max: 50 }} // Faster vertical fall to ensure quick exit
    />
  );
};

export default ConfettiComponent;
