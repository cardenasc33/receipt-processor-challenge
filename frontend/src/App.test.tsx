import React from 'react';
import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import App from './App';
import fetchMock from 'jest-fetch-mock';

// 👇 MOCK react-confetti to avoid canvas crash
jest.mock('react-confetti', () => () => <div data-testid="mock-confetti" />);

beforeEach(() => {
  fetchMock.resetMocks();
  jest.useFakeTimers();
  jest.spyOn(window, 'alert').mockImplementation(() => {});
});

afterEach(() => {
  jest.runOnlyPendingTimers();
  jest.useRealTimers();
});

test('renders receipt form title', () => {
  render(<App />);
  expect(screen.getByText(/Submit Receipt Details/i)).toBeInTheDocument();
});

test('adds item to the list', async () => {
  render(<App />);
  await userEvent.type(screen.getByLabelText(/Item Name/i), 'Test Item');
  await userEvent.type(screen.getByLabelText(/Item Price/i), '9.99');
  await userEvent.click(screen.getByText(/Add Item/i));
  expect(screen.getByText(/Test Item: \$9.99/i)).toBeInTheDocument();
});

test('shows alert on empty receipt details', async () => {
  render(<App />);
  await userEvent.click(screen.getByRole('button', { name: /Submit Receipt/i }));
  expect(window.alert).toHaveBeenCalledWith(expect.stringMatching(/fill in all the receipt details/i));
});

test('submits receipt and fetches points', async () => {
  fetchMock.mockResponses(
    [JSON.stringify({ id: 'abc123' }), { status: 200 }],
    [JSON.stringify({ points: '500' }), { status: 200 }]
  );

  render(<App />);
  await userEvent.type(screen.getByLabelText(/Retailer/i), 'StoreX');
  await userEvent.type(screen.getByLabelText(/Purchase Date/i), '2023-10-10');
  await userEvent.type(screen.getByLabelText(/Purchase Time/i), '12:00');
  await userEvent.type(screen.getByLabelText(/Item Name/i), 'Apple');
  await userEvent.type(screen.getByLabelText(/Item Price/i), '2.00');
  await userEvent.click(screen.getByText(/Add Item/i));
  await userEvent.click(screen.getByRole('button', { name: /Submit Receipt/i }));

  expect(await screen.findByText(/Receipt ID: abc123/i)).toBeInTheDocument();
  expect(await screen.findByText(/Points: 500/i)).toBeInTheDocument();
});

test('displays modal and triggers confetti after successful receipt submission', async () => {
  fetchMock.mockResponses(
    [JSON.stringify({ id: 'xyz789' }), { status: 200 }],
    [JSON.stringify({ points: '999' }), { status: 200 }]
  );

  render(<App />);
  await userEvent.type(screen.getByLabelText(/Retailer/i), 'TestMart');
  await userEvent.type(screen.getByLabelText(/Purchase Date/i), '2023-12-25');
  await userEvent.type(screen.getByLabelText(/Purchase Time/i), '18:30');
  await userEvent.type(screen.getByLabelText(/Item Name/i), 'Candy');
  await userEvent.type(screen.getByLabelText(/Item Price/i), '3.50');
  await userEvent.click(screen.getByText(/Add Item/i));
  await userEvent.click(screen.getByRole('button', { name: /Submit Receipt/i }));

  expect(await screen.findByText(/Receipt ID: xyz789/i)).toBeInTheDocument();
  expect(screen.getByText(/Retailer: TestMart/i)).toBeInTheDocument();
  expect(screen.getByText(/Points: 999/i)).toBeInTheDocument();
  expect(screen.getByTestId('mock-confetti')).toBeInTheDocument();
});
