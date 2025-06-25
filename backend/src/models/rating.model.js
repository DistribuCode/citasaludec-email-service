import mongoose from 'mongoose';

const ratingSchema = new mongoose.Schema({
  doctorId: { type: String, required: true },
  patientId: { type: String, required: true },
  score: { type: Number, required: true, min: 1, max: 5 },
  comment: { type: String }
}, { timestamps: true });

const Rating = mongoose.model('Rating', ratingSchema);
export default Rating;