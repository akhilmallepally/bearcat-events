from rest_framework import serializers
from .models import bearcatE

class bearcatE_serializer(serializers.ModelSerializer):
    class Meta:
        model = bearcatE
        fields = ('id', 'title', 'description', 'completed')
