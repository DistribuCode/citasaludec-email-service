import express from 'express';
import { createRating, getAllRatings } from '../controllers/rating.controller.js';

const router = express.Router();
router.post('/', createRating);
router.get('/', getAllRatings);

export default router;