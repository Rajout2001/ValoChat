import { useState } from "react";
import { validateEmail } from "../utils";
import { subscribeToNewsletter } from "../services/NewsletterService";

const Footer: React.FC = () => {
  const [newsletterEmail, setNewsletterEmail] = useState('');

  const handleNewsletterSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!validateEmail(newsletterEmail)) {
      return;
    }
    
    const data = await subscribeToNewsletter(newsletterEmail); 
    if (data?.code === '201') {
      alert("You have successfully subscribed to our newsletter!");
      setNewsletterEmail('');
    } else {
      alert("You have already subscribed to our newsletter!");
    }
    
  }

  return (
    <footer className="bg-gray-900 text-white py-9">
      <div className="container mx-auto flex flex-col lg:flex-row justify-between items-center">
        <div className="text-lg font-bold mb-4 lg:mb-0">Valorant Matchs Results</div>
          <div className="mb-4 lg:mb-0">
            <p className="text-sm">Subscribe to our newsletter for weekly updates</p>
            <form className="flex mt-2" onSubmit={handleNewsletterSubmit}>
              <input
                className="bg-gray-800 text-white px-4 py-2 rounded-l focus:outline-none"
                type="email"
                placeholder="Your email"
                value={newsletterEmail}
                onChange={(e) => setNewsletterEmail(e.target.value)}
              />
              <button
                type="submit"
                className="bg-blue-500 hover:bg-blue-600 px-4 py-2 rounded-r focus:outline-none"
              >
                Subscribe
              </button>
            </form>
          </div>
          <div className="text-sm ml-0 lg:ml-6 mt-6 lg:mt-0">
            <p>&copy; {new Date().getFullYear()} All rights reserved</p>
            <p>Created by: </p>  
            <p className="font-satoshi">BOUZOURINE Hichem & Rajith Ravindran</p>
          </div>
      </div>
    </footer>
  );
};

export default Footer;
