import Rating from '../models/rating.model.js';

export const createRating = async (req, res) => {
  try {
    const newRating = new Rating(req.body);
    const savedRating = await newRating.save();
    res.status(201).json(savedRating);
  } catch (error) {
    res.status(500).json({ error: error.message });
  }
};

export const getAllRatings = async (req, res) => {
  try {
    const ratings = await Rating.find();
    res.json(ratings);
  } catch (error) {
    res.status(500).json({ error: error.message });
  }
};