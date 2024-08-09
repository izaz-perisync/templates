import { FormspreeProvider } from '@formspree/react';
function App({ Component, pageProps }) {
  return (
    <FormspreeProvider project="2535331811941154503">
      <Component {...pageProps} />
    </FormspreeProvider>
  );
}
export default App;
