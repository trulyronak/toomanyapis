from flask_api import FlaskAPI
import os

app = FlaskAPI("Optic")

@app.route('/')
def index():
	return "hi world"

app.run(port=os.environ['OPTIC_API_PORT'] if 'OPTIC_API_PORT' in os.environ else 4003)
