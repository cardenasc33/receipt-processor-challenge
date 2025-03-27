import React, { useState } from 'react';
import './App.css'; // Import the CSS file for styling
import Modal from './Modal'; 
import './Modal.css'; 

interface Item {
  Description: string;
  Price: string;
}


const App: React.FC = () => {
  const [itemName, setItemName] = useState('');
  const [itemPrice, setItemPrice] = useState('');
  const [retailer, setRetailer] = useState('');
  const [purchaseDate, setPurchaseDate] = useState('');
  const [purchaseTime, setPurchaseTime] = useState('');
  const [items, setItems] = useState<Item[]>([]);
  const [receiptID, setReceiptID] = useState('')
  const [points, setPoints] = useState('')

  // Modal used to show submission of receipt
  const [isModalOpen, setIsModalOpen] = useState(false);

  const openModal = () => setIsModalOpen(true);
  const closeModal = () => {
    // Clear form after submission
    setRetailer('');
    setPurchaseDate('');
    setPurchaseTime('');
    setItems([]);
    setIsModalOpen(false);
  }
  

  const totalPrice = items.reduce((total, item) => total + parseFloat(item.Price), 0);
  console.log("Total Price: " , totalPrice)


  const data = {
    retailer: retailer,
    purchaseDate: purchaseDate,
    purchaseTime: purchaseTime,
    total: totalPrice.toFixed(2),
    items: items,
  };

  const handleAddItem = () => {
    const price = parseFloat(itemPrice);
    if (itemName && !isNaN(price) && price > 0) {
      const newItem = { Description: itemName, Price: price.toString() };
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
  

    try {
    

      const response = await fetch('http://localhost:8000/receipts/process', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data)
      });

      if (!response.ok) {
        throw new Error('Failed to submit receipt');
      }

      const responseData = await response.json();
      console.log('Receipt Submitted:', responseData);

      setReceiptID(responseData.ReceiptID)
      setPoints(responseData.Points)
      openModal() 

   
      
    } catch (error) {
      console.error('Error submitting receipt:', error);
      alert('Error submitting receipt. Please try again later.');
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

      <div>
      <Modal isOpen={isModalOpen} onClose={closeModal} title="Receipt Submitted!">
        <p>{"\nRetailer: " + retailer}</p>
        <p>{"\nReceipt ID: " + receiptID}</p>
        <p>{"\nPoints: " + points }</p>
      </Modal>
    </div>


    </div>
  );
};

export default App;
