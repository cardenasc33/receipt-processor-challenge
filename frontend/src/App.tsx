import React, { useState, useEffect } from 'react';
import './App.css';
import Modal from './Modal';
import './Modal.css';
import ConfettiComponent from './ConfettiComponent';

interface Item {
  Description: string;
  Price: string;
}

interface PostResponse {
  ReceiptID: string;
}

interface GetPointsResponse {
  Points: string;
}

const hostIp = process.env.DOCKER_HOST_IP || '127.0.0.1';
const port = process.env.PORT || '8080'; // Ensure it's a string
const backendUrl = `http://${hostIp}:${port}`;

const App: React.FC = () => {
  const [itemName, setItemName] = useState('');
  const [itemPrice, setItemPrice] = useState('');
  const [retailer, setRetailer] = useState('');
  const [purchaseDate, setPurchaseDate] = useState('');
  const [purchaseTime, setPurchaseTime] = useState('');
  const [items, setItems] = useState<Item[]>([]);
  const [receiptID, setReceiptID] = useState('');
  const [points, setPoints] = useState('');
  const [runConfetti, setRunConfetti] = useState(false);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [dimensions, setDimensions] = useState({ width: 0, height: 0 });

  useEffect(() => {
    setDimensions({
      width: window.innerWidth,
      height: window.innerHeight,
    });
  }, []);

  const showConfetti = () => {
    setRunConfetti(true);
    setTimeout(() => {
      setRunConfetti(false);
    }, 6000);
  };

  useEffect(() => {
    if (receiptID !== '') {
      getReceipt();
      openModal();
    }
  }, [receiptID]);

  const openModal = () => {
    setIsModalOpen(true);
    showConfetti();
  };

  const closeModal = () => {
    setRetailer('');
    setPurchaseDate('');
    setPurchaseTime('');
    setItems([]);
    setIsModalOpen(false);
  };

  const totalPrice = items.reduce((total, item) => total + parseFloat(item.Price), 0);

  const handleAddItem = () => {
    const price = parseFloat(itemPrice);
    if (itemName && !isNaN(price) && price > 0) {
      const newItem: Item = { Description: itemName, Price: price.toString() };
      setItems((prevItems) => [...prevItems, newItem]);
      setItemName('');
      setItemPrice('');
    } else {
      alert('Please enter a valid item name and price.');
    }
  };

  const handleSubmitReceipt = async () => {
    if (!retailer || !purchaseDate || !purchaseTime) {
      alert('Please fill in all the receipt details (Retailer, Date, Time).');
      return;
    }

    if (items.length === 0) {
      alert('Please add at least one item to the receipt.');
      return;
    }

    const data = {
      retailer,
      purchaseDate,
      purchaseTime,
      total: totalPrice.toFixed(2),
      items,
    };

    // POST Request
    try {
      const postResponse = await fetch(`${backendUrl}/receipts/process`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      });

      if (!postResponse.ok) {
        throw new Error('Failed to submit receipt');
      }

      const postData: PostResponse = await postResponse.json();
      console.log('Receipt Submitted:', postData);
      setReceiptID(postData.ReceiptID);
    } catch (error) {
      console.error('Error submitting receipt:', error);
      alert('Error submitting receipt. Please try again later.');
    }
  };

  // GET Request
  const getReceipt = async () => {
    try {
      const getResponse = await fetch(`${backendUrl}/receipts/${receiptID}/points`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (!getResponse.ok) {
        throw new Error('Failed to retrieve receipt');
      }

      const getData: GetPointsResponse = await getResponse.json();
      console.log('Receipt Retrieved:', getData);
      setPoints(getData.Points);
    } catch (error) {
      console.error('Error retrieving receipt:', error);
      alert('Error retrieving receipt points.');
    }
  };

  return (
    <div className="form-container">
      <h2 className="form-title">Submit Receipt Details</h2>

      <div className="form-section">
        <label htmlFor="retailer" className="form-label">Retailer:</label>
        <input
          id="retailer"
          type="text"
          value={retailer}
          onChange={(e) => setRetailer(e.target.value)}
          placeholder="Enter retailer name"
          className="form-input"
        />
      </div>

      <div className="form-section">
        <label htmlFor="purchaseDate" className="form-label">Purchase Date:</label>
        <input
          id="purchaseDate"
          type="date"
          value={purchaseDate}
          onChange={(e) => setPurchaseDate(e.target.value)}
          className="form-input"
        />
      </div>

      <div className="form-section">
        <label htmlFor="purchaseTime" className="form-label">Purchase Time:</label>
        <input
          id="purchaseTime"
          type="time"
          value={purchaseTime}
          onChange={(e) => setPurchaseTime(e.target.value)}
          className="form-input"
        />
      </div>

      <div className="form-section">
        <label htmlFor="itemName" className="form-label">Item Name:</label>
        <input
          id="itemName"
          type="text"
          value={itemName}
          onChange={(e) => setItemName(e.target.value)}
          placeholder="Enter item name"
          className="form-input"
        />
      </div>

      <div className="form-section">
        <label htmlFor="itemPrice" className="form-label">Item Price:</label>
        <input
          id="itemPrice"
          type="number"
          value={itemPrice}
          onChange={(e) => setItemPrice(e.target.value)}
          placeholder="Enter item price"
          className="form-input"
        />
      </div>

      <button type="button" onClick={handleAddItem} className="btn-add-item">Add Item</button>

      <h3 className="items-list-title">Items List:</h3>
      <ul className="items-list">
        {items.map((item, index) => (
          <li key={index} className="item">
            {item.Description}: ${parseFloat(item.Price).toFixed(2)}
          </li>
        ))}
      </ul>

      {items.length > 0 && (
        <div className="total-container">
          <h3 className="total-title">Total: ${totalPrice.toFixed(2)}</h3>
        </div>
      )}

      <button type="button" onClick={handleSubmitReceipt} className="btn-submit-receipt">
        Submit Receipt
      </button>

      <ConfettiComponent
        width={dimensions.width}
        height={dimensions.height}
        runConfetti={runConfetti}
        gravity={1}
        numberOfPieces={175}
      />

      <Modal isOpen={isModalOpen} onClose={closeModal} title="Receipt Submitted!">
        <p>{"\nRetailer: " + retailer}</p>
        <p>{"\nReceipt ID: " + receiptID}</p>
        <p>{"\nPoints: " + points}</p>
      </Modal>
    </div>
  );
};

export default App;
