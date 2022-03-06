from rest_framework import serializers
from .models import EventModel

class bearcatESerializer(serializers.ModelSerializer):
	class Meta:
		model = EventModel
		fields = ('id', 'title', 'description', 'location')