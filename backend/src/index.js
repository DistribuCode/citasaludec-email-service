import express from 'express';
import cors from 'cors';
import dotenv from 'dotenv';
import { connectDB } from './config/db.js';
import ratingRoutes from './routes/rating.routes.js';

dotenv.config();
const app = express();
app.use(cors());
app.use(express.json());

connectDB();

app.use('/api/ratings', ratingRoutes);

const PORT = process.env.PORT || 5017;
app.listen(PORT, () => {
  console.log(`🚀 Rating Service running on port ${PORT}`);
});