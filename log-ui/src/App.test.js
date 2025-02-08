import { render, screen } from '@testing-library/react';
import App from './App';

test('renders learn react link', () => {
  render(<App />);
  const linkElement = screen.getByText(/Tell the endpoint the current time/i);
  expect(linkElement).toBeInTheDocument();
});

