const bcrypt = require('bcryptjs');
const jwt = require('jsonwebtoken');

const User = require('../models/user.models');

exports.register = async (req, res) => {
  try {

    const {
      name,
      email,
      password
    } = req.body;

    const existing =
      await User.findByEmail(email);

    if (existing) {
      return res.status(400).json({
        message: 'Email already exists'
      });
    }

    const hashedPassword =
      await bcrypt.hash(password, 10);

    const user =
      await User.createUser(
        name,
        email,
        hashedPassword,
        'user'
      );

    res.status(201).json({
      message: 'User created',
      data: user
    });

  } catch (err) {

    res.status(500).json({
      message: err.message
    });

  }
};

exports.login = async (req, res) => {

  try {

    const {
      email,
      password
    } = req.body;

    const user =
      await User.findByEmail(email);

    if (!user) {

      return res.status(401).json({
        message: 'Invalid credentials'
      });

    }

    const valid =
      await bcrypt.compare(
        password,
        user.password
      );

    if (!valid) {

      return res.status(401).json({
        message: 'Invalid credentials'
      });

    }

    const token = jwt.sign(
      {
        id: user.id,
        role: user.role,
        email: user.email
      },
      process.env.JWT_SECRET,
      {
        expiresIn: '1d'
      }
    );

    res.json({
      token,
      user: {
        id: user.id,
        name: user.name,
        email: user.email,
        role: user.role
      }
    });

  } catch (err) {

    res.status(500).json({
      message: err.message
    });

  }
};