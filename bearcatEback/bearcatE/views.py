from django.shortcuts import render
from rest_framework import viewsets
from .serializers import bearcatE_serializer
from .models import bearcatE


# Create your views here.

class bearcatE_view(viewsets.ModelViewSet):
    serializer_class = bearcatE_serializer
    queryset = bearcatE.objects.all()

