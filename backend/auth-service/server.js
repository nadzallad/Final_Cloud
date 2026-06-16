require('dotenv').config();

const express = require('express');
const cors = require('cors');

const authRoutes =
  require('./routes/auth.routes');

const app = express();

app.use(cors());
app.use(express.json());

app.use('/auth', authRoutes);

app.get('/', (req, res) => {
  res.json({
    service: 'Auth Service',
    status: 'Running'
  });
});

app.listen(process.env.PORT, () => {
  console.log(
    `Auth Service running on port ${process.env.PORT}`
  );
});